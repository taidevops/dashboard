interface Window {
  define: (name: string, deps: string[], definitionFn: () => any) => void;

  System: {
    // @ts-ignore
    import: (path) => Promise<any>;
  };
}
