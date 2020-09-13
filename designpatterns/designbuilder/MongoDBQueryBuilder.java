public class MongoDBQueryBuilder implements QueryBuilder{
	private MongoDBQuery query = new MongoDBQuery();

	public void setFrom(String from){
		query.setFrom(from);
	}

	public void setWhere(String where){
		query.setWhere(where);
	}

	public Query getQuery(){
		return query;
	}
}
