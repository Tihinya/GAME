export function isEvent(key: string, value: any): boolean {
	return key.substring(0, 2) === "on" && value instanceof Object
}

export function isProps(key: string): boolean {
	return key !== "children"
}

export function propertyHasChanged(
	oldKey: string,
	value: string,
	newProps: ElementProps
): boolean {
	return !newProps.hasOwnProperty(oldKey) || newProps[oldKey] !== value
}

export function isArraysEqual<T>(a: T[], b: T[]): boolean {
	return (
		a.length === b.length &&
		a.some((value: any, i: number) => !Object.is(value, b[i]))
	)
}
