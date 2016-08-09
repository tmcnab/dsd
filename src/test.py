import options
import unittest


def _build_suite():
    suite = unittest.TestSuite()
    suite.addTest(options.OptionsTests())
    return suite

def test():
    suite = _build_suite()
    runner = unittest.TextTestRunner()
    runner.run(suite)
