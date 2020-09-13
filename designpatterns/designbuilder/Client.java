public class Client{
	public static void main(String args[]){
		QueryBuilderDirector queryBuilderDirector = new QueryBuilderDirector();

		String from = "clientdb";
		String where = "name=\"robert\"";

		MongoDBQueryBuilder mongoDBQueryBuilder = new MongoDBQueryBuilder();
		Query query = queryBuilderDirector.buildQuery(from, where, mongoDBQueryBuilder);
		query.execute();


		NoSQLQueryBuilder noSQLQueryBuilder = new NoSQLQueryBuilder();
		query = queryBuilderDirector.buildQuery(from, where, noSQLQueryBuilder);
		query.execute();
	}
}
