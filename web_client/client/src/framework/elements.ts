export function CreateElement(
    parent: HTMLElement, 
    type: string, 
    classes: string[], 
    onMount: (element: HTMLElement) => void
): HTMLElement {
    const element = document.createElement(type);
    classes.forEach((value) => {element.classList.add(value);});
    parent.appendChild(element);
    onMount(element);
    return element;
}