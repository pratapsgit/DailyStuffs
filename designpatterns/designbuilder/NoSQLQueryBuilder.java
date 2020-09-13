public class NoSQLQueryBuilder implements QueryBuilder{
	private NoSQLQuery query = new NoSQLQuery();

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
