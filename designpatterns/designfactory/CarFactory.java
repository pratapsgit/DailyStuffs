public class CarFactory extends TransportFactory{
	public Transport create() {
		return new Car();
	}
}
