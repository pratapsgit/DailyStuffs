public class SQLQueryBuilder implements QueryBuilder{
	private SQLQuery query = new SQLQuery();

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
