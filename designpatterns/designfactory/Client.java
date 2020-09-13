public class Client{
	public static void main(String args[]){
		TransportFactory transportFactory = new CarFactory();
		Transport transport = transportFactory.create();
		System.out.println(transport.drive());


		transportFactory = new ShipFactory();
		transport = transportFactory.create();
		System.out.println(transport.drive());
	}
}
