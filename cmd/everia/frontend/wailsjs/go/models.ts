export namespace models {
	
	export class AppCatalog {
	    name: string;
	    version: string;
	    image_url: string;
	
	    static createFrom(source: any = {}) {
	        return new AppCatalog(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.version = source["version"];
	        this.image_url = source["image_url"];
	    }
	}
	export class WebsiteCatalog {
	    name: string;
	    url: string;
	    version: string;
	    image_url: string;
	
	    static createFrom(source: any = {}) {
	        return new WebsiteCatalog(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.url = source["url"];
	        this.version = source["version"];
	        this.image_url = source["image_url"];
	    }
	}
	export class Catalog {
	    name: string;
	    authors: string[];
	    last_modified: string;
	    apps: AppCatalog[];
	    websites: WebsiteCatalog[];
	
	    static createFrom(source: any = {}) {
	        return new Catalog(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.authors = source["authors"];
	        this.last_modified = source["last_modified"];
	        this.apps = this.convertValues(source["apps"], AppCatalog);
	        this.websites = this.convertValues(source["websites"], WebsiteCatalog);
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
	export class PasswordEntry {
	    site: string;
	    username: string;
	    password: string;
	    notes: string;
	
	    static createFrom(source: any = {}) {
	        return new PasswordEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.site = source["site"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.notes = source["notes"];
	    }
	}

}

export namespace note {
	
	export class Info {
	    title: string;
	    lastEditTime: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.lastEditTime = source["lastEditTime"];
	    }
	}

}

