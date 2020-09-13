public class Client {
	public static void apply(ThemeFactory themeFactory){
		String themeString = String.format("We have a theme with %s %s %s", themeFactory.createButton().getDetails(),
				themeFactory.createDialog().getDetails(),
				themeFactory.createTitleBar().getDetails());
		System.out.println(themeString);
	}

	public static void main(String args[]) {
		ThemeFactory themeFactory = new DarkThemeFactory();
		apply(themeFactory);

		themeFactory = new LightThemeFactory();
		apply(themeFactory);
	}
}
