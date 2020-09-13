public class Button implements Widget{
	private String color = "None";
	public Button(){
	}

	public Button(String color){
		this.color = color;
	}

	public String getDetails() {
		return this.color+ " " + this.getClass().getName();
	}
}
