import os
import pickle
import threading
import uuid


class AddObjectThread(threading.Thread):

    def __init__(self, cluster, item, id=''):
        """Initialize AddObjectThread.

        Create a uniquely named "lock" file which has the dictionary pickled
        into it. The thread will remove the file when the data has been
        permanently persisted. If the file is present on starting DSD, that
        means that DSD crashed before the data could be integrated. So ...
        recover it!
        """
        super().__init__()
        self._cluster = cluster
        self.item = item
        self.id = uuid.uuid4()
        with open(self.filename, 'wb') as file:
            pickle.dump(item, file)

    @property
    def filename(self):
        return self.id.hex + '.pickle'

    def run(self):
        # If the item already exists in the store then we won't even need to
        # let the cluster know that it has been inserted.
        if self._item_exists:
            os.remove(self.filename)
            return

        # If the item doesn't exist, append to the data log and let any indices
        # know they need to consider it.
        self._cluster.insert(self.item)

    @property
    def _item_exists(self):
        return False


class Store(object):

    def __init__(self, cluster):
        self._cluster = cluster

    def add_object(self, dict):
        AddObjectThread(self._cluster, dict).start()

    def integrity_check(self):
        pass
