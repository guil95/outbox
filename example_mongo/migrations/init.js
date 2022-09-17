db = db.getSiblingDB('app_mongo_db');

db.createCollection('outbox');