const baseHref = 'api/v1';

export enum Resource {
  job = 'job',
  cronJob = 'cronjob',
  pod = 'pod'
}

export enum Utility {
  shell = 'shell',
}

class ResourceEndpoint {
  constructor(private readonly resource_: Resource, private readonly namespaced_ = false) {}

  list(): string {
    return `${baseHref}/${this.resource_}${this.namespaced_ ? '/:namespace' : ''}`;
  }

  detail(): string {
    return `${baseHref}/${this.resource_}${this.namespaced_ ? '/:namespace' : ''}/:name`;
  }
}

class UtilityEndpoint {
  constructor(private readonly utility_: Utility) {}

  shell(namespace: string, resourceName: string): string {
    return `${baseHref}/${Resource.pod}/${namespace}/${resourceName}/${this.utility_}`;
  }
}

export class EndpointManager {
  static resource(resource: Resource, namespaced?: boolean): ResourceEndpoint {
    return new ResourceEndpoint(resource, namespaced);
  }

  static utility(utility: Utility): UtilityEndpoint {
    return new UtilityEndpoint(utility);
  }
}