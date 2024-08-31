export namespace models {
	
	export class Brand {
	    id?: number;
	    name?: string;
	    // Go type: time
	    created_at?: any;
	    // Go type: time
	    updated_at?: any;
	
	    static createFrom(source: any = {}) {
	        return new Brand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Item {
	    id?: number;
	    item_code?: string;
	    item_name?: string;
	    item_description?: string;
	    item_status?: boolean;
	    quantity?: number;
	    category_id?: number;
	    brand_id?: number;
	    discount?: number;
	    // Go type: time
	    created_at?: any;
	    // Go type: time
	    updated_at?: any;
	    brand?: Brand;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.item_code = source["item_code"];
	        this.item_name = source["item_name"];
	        this.item_description = source["item_description"];
	        this.item_status = source["item_status"];
	        this.quantity = source["quantity"];
	        this.category_id = source["category_id"];
	        this.brand_id = source["brand_id"];
	        this.discount = source["discount"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.brand = this.convertValues(source["brand"], Brand);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

