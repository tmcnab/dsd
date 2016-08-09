import socket
import sys
import unittest


class ClusterServer(object):
    """Clustering server.

    The clustering server performs two actions:
    1. coordinates logical timestamps between members of the cluster; and
    2. replicates objects recieved by this instance to other instances.
    """

    def __init__(self, options):
        """Initialize a new instance of the ClusterServer."""
        self._options = options

    def check_replication(self, set, subsets):
        """Check the status of replication of objects and subsets over the cluster."""
        pass

    def inserted(self, item):
        """Notify members of object insertion."""
        pass


class ClusterServerTests(unittest.TestCase):
    """Unit tests for the ClusterServer."""

    pass
