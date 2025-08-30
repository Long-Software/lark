export const mapWithSeparator = <A, B>(items: A[], sep: B): (A | B)[] => {
  const result: (A | B)[] = []

  items.forEach((item, index) => {
    result.push(item)
    if (index < items.length - 1) {
      result.push(sep)
    }
  })

  return result
}
