public class TitleBar implements Widget{
	private String color = "None";
	public TitleBar(){
	}

	public TitleBar(String color){
		this.color = color;
	}

	public String getDetails() {
		return this.color+ " " + this.getClass().getName();
	}
}
