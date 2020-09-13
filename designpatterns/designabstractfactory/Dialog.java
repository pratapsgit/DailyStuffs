public class Dialog implements Widget{
	private String color = "None";
	public Dialog(){
	}

	public Dialog(String color){
		this.color = color;
	}

	public String getDetails() {
		return this.color+ " " + this.getClass().getName();
	}
}
