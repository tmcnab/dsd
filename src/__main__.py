from cluster import Cluster
from store import Store

cluster = Cluster()
store = Store(cluster)
o = {'name': 'Alice'}
store.add_object(o)
