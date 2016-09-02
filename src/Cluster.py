import socket
import sys
import unittest


class Cluster(object):
    """Clustering server.

    The clustering server performs two actions:
    1. coordinates logical timestamps between members of the cluster; and
    2. replicates objects recieved by this instance to other instances.
    """

    def check_replication(self, set, subsets):
        """Check the status of replication of objects and subsets over the cluster."""
        pass

    def insert(self, item):
        """Notify members of object insertion."""
        print('Letting cluster know that item was inserted')


class ClusterServerTests(unittest.TestCase):
    """Unit tests for the ClusterServer."""

    pass
