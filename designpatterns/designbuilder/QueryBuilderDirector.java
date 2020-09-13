public class QueryBuilderDirector{
	public Query buildQuery(String from, String where, QueryBuilder builder){
		builder.setFrom(from);
		builder.setWhere(where);
		return builder.getQuery();
	}
}
