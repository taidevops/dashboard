
export class GroupedResourceList {

  private readonly items_: {[id: string]: number} = {};

  shouldShowZeroState(): boolean {
    let totalItems = 0;
    const ids = Object.keys(this.items_);
    ids.forEach(id => (totalItems += this.items_[id]));
    return totalItems === 0 && ids.length > 0;
  }

  
}