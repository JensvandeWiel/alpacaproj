export type Errors = Record<string, string>;
export type ErrorBag = Record<string, Errors>;
export type StandardPageProps<T extends Record<string, unknown>> = T & {
    "errors": Errors & ErrorBag
}