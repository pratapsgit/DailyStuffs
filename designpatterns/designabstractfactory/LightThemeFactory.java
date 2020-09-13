public class LightThemeFactory implements ThemeFactory {
	public Widget createButton() {
		return new Button("White");
	}

	public Widget createDialog() {
		return new Dialog("White");
	}

	public Widget createTitleBar() {
		return new TitleBar("White");
	}
}
