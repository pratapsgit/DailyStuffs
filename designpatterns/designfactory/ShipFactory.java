public class ShipFactory extends TransportFactory {
	public Transport create(){
		return new Ship();
	}
}
