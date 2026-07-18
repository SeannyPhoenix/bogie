export namespace JSX {
	export type Element = HTMLElement | Text;

	export type GlobalAttributes = {
		accessKey?: string;
		autofocus?: boolean;
		class?: string;
		contentEditable?: boolean;
		data?: { [key: string]: any };
		dir?: string;
		draggable?: boolean;
		enterKeyHint?: string;
		exportparts?: string;
		hidden?: boolean;
		id?: string;
		inert?: boolean;
		inputMode?: string;
		itemID?: string;
		itemProp?: string;
		itemRef?: string;
		itemScope?: boolean;
		itemType?: string;
		lang?: string;
	}

	export type IntrinsicElements = {
		[a: string]: any;
	}
	export interface ElementChildrenAttribute {
		children: unknown;
	}
	export type FComponent<P = {}> = (props: P) => Element;
}
