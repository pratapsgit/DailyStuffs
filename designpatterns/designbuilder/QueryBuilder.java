public interface QueryBuilder{
	public void setFrom(String from);
	public void setWhere(String where);
	public Query getQuery();
}
