// usecase:
//mongo -u root -p 12345 --authenticationDatabase admin  192.168.2.91:27017/gamemain ./dbscripts/create_db.js > ./dbscripts/create.log
//mongo mongodb://root:12345@192.168.2.91:27017/gamemain ./dbscripts/create_db.js > ./dbscripts/create.log
var s = db.getMongo().startSession();
var user_collection = "game_user";
s.startTransaction();
print("==========================> StartTime"+Date("<YYYY-mm-ddTHH:MM:ssZ>"));
try {
    db.createCollection(user_collection,{autoIndexId:1});
    db[user_collection].createIndex({"userid":1},{unique:true});
    // db[user_collection].insert({userid:NumberLong(1)});
    // db[user_collection].drop({userid:NumberLong(1)});
    s.commitTransaction()
} catch(e) {
  print(e);
  s.abortTransaction()
}
print("==========================> EndTime"+Date("<YYYY-mm-ddTHH:MM:ssZ>"));
print("Transaction completed...");