export namespace main {
	
	export class FileList {
	    Extension: string;
	    Name: string;
	    AbsolutePath: string;
	
	    static createFrom(source: any = {}) {
	        return new FileList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Extension = source["Extension"];
	        this.Name = source["Name"];
	        this.AbsolutePath = source["AbsolutePath"];
	    }
	}
	export class Result {
	    Files: FileList[];
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Files = this.convertValues(source["Files"], FileList);
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

