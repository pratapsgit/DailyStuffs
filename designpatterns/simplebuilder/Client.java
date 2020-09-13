public class Client{
	public static void main(String args[]){
		SQLQuery sqlQuery = new SQLQuery();
		sqlQuery.setFrom("sqldb");
		sqlQuery.setWhere("name=\"*duster*\"");
		sqlQuery.execute();


		MongoDBQuery mongoDBQuery = new MongoDBQuery();
		mongoDBQuery.setFrom("sqldb");
		mongoDBQuery.setWhere("name=\"*duster*\"");
		mongoDBQuery.execute();
	}
}
