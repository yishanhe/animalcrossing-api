import pymongo
import json

MONGO_URL = "mongodb://localhost:27017"
DATABASE_NAME = "AnimalCrossingDB"
COLLECTION_NAME = "fishes"
client = pymongo.MongoClient(MONGO_URL)
database = client[DATABASE_NAME]


JSON_PATH = "/Users/syi/Github/google-sheets-to-json/out/"

with open(JSON_PATH+"creatures.json", 'r') as json_file:

    data = json_file.read()
    obj = json.loads(data)
    fishes_jsons = filter(lambda x: x["sourceSheet"] == "Fish", obj)
    bugs_jsons = filter(lambda x: x["sourceSheet"] == "Bugs", obj)

    fishes = database["fishes"]
    fishes.drop()
    fishes.insert_many(fishes_jsons)

    bugs = database["bugs"]
    bugs.drop()
    bugs.insert_many(bugs_jsons)


with open(JSON_PATH+"construction.json", 'r') as json_file:
    data = json_file.read()
    obj = json.loads(data)
    coll = database["constructions"]
    coll.drop()
    coll.insert_many(obj)


with open(JSON_PATH+"items.json", 'r') as json_file:
    data = json_file.read()
    obj = json.loads(data)
    coll = database["items"]
    coll.drop()
    coll.insert_many(obj)

with open(JSON_PATH+"nookMiles.json", 'r') as json_file:
    data = json_file.read()
    obj = json.loads(data)
    coll = database["nook_miles"]
    coll.drop()
    coll.insert_many(obj)

with open(JSON_PATH+"recipes.json", 'r') as json_file:
    data = json_file.read()
    obj = json.loads(data)
    coll = database["recipes"]
    coll.drop()
    coll.insert_many(obj)

with open(JSON_PATH+"villagers.json", 'r') as json_file:
    data = json_file.read()
    obj = json.loads(data)
    coll = database["villagers"]
    coll.drop()
    coll.insert_many(obj)
