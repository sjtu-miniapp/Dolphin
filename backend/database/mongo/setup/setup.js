db.auth('admin', 'admin')
db = db.getSiblingDB('dolphin');

db.createUser(
    {
        user: "root",
        pwd: "610878",
        roles: [
            { role: "dbOwner", db: "dolphin"}
        ]
    }
);

db.createCollection("task");
db.task.createIndex({"task_id": 1});
