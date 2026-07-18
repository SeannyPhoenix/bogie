import type { HTMLIntrinsicElements } from "./types";

export interface ElementFactory {
  createElement(localName: string): HTMLElement;
  createTextNode(data: string): Text;
}

let elementFactory: ElementFactory | null = null;

if ("document" in globalThis) {
  elementFactory = {
    createElement(localName: string): HTMLElement {
      return document.createElement(localName);
    },
    createTextNode(data: string): Text {
      return document.createTextNode(data);
    },
  };
}

export function setElementFactory(factory: ElementFactory): void {
  elementFactory = factory;
}

export function getElementFactory(): ElementFactory {
  if (!elementFactory) {
    throw new Error("Element factory not set");
  }
  return elementFactory;
}

export const Fragment = Symbol("Fragment");

export namespace JSX {
  export type Element = HTMLElement | Text | DocumentFragment;
  export type IntrinsicElements = HTMLIntrinsicElements;
  export interface ElementChildrenAttribute {
    children: unknown;
  }
  export type FComponent<P = {}> = (props: P) => Element;
}

export function jsx(type: any, props: any, key?: string | number): JSX.Element {
  if (typeof type === "string") {
    const factory = getElementFactory();
    const element = factory.createElement(type);
    for (const [name, value] of Object.entries(props)) {
      if (name === "children") {
        appendChildren(element, value);
      } else if (name.startsWith("on") && typeof value === "function") {
        const eventName = name
          .slice(2)
          .toLowerCase() as keyof HTMLElementEventMap;
        element.addEventListener(eventName, value as EventListener);
      } else if (value !== undefined) {
        element.setAttribute(name, String(value));
      }
    }
    return element;
  } else if (typeof type === "function") {
    return type(props);
  } else if (type === Fragment) {
    const fragment = document.createDocumentFragment();
    appendChildren(fragment, props.children);
    return fragment;
  } else {
    throw new Error(`Unsupported JSX type: ${type}`);
  }
}

export const jsxs = jsx;

function appendChildren(
  parent: HTMLElement | DocumentFragment,
  children: any,
): void {
  if (Array.isArray(children)) {
    for (const child of children) {
      appendChildren(parent, child);
    }
  } else if (
    children instanceof HTMLElement ||
    children instanceof Text ||
    children instanceof DocumentFragment
  ) {
    parent.appendChild(children);
  } else if (children != null && children !== false && children !== undefined) {
    const factory = getElementFactory();
    const textNode = factory.createTextNode(String(children));
    parent.appendChild(textNode);
  }
}
