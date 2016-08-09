import unittest


class Options(object):
    """docstring for Options."""

    def __init__(self, dict):
        """Initialize a new instance of the Options type."""
        self._raw = dict

    @property
    def peers(self):
        """All IPv6 peer addresses."""
        return self._raw['peers'] or []


class OptionsTests(unittest.TestCase):
    """Unit tests for the Options type."""

    def test_peers(self):
        """Test that peers exists."""
        options = Options({})
        self.assertIn('peers', options)
