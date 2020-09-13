public class NoSQLQuery implements Query{
	private String from;
	private String where;

	public void setFrom(String from){
		this.from = from;
	}

	public void setWhere(String where){
		this.where = where;
	}

	public void execute(){
		System.out.printf("Executing No SQL query \"select * from %s where %s\"\n", this.from, this.where);
	}
}
