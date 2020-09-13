public class DarkThemeFactory implements ThemeFactory {
	public Widget createButton() {
		return new Button("Black");
	}

	public Widget createDialog() {
		return new Dialog("Black");
	}

	public Widget createTitleBar() {
		return new TitleBar("Black");
	}
}
