export namespace models {
	
	export class ClipboardItem {
	    content: string;
	    createdAt: string;
	
	    static createFrom(source: any = {}) {
	        return new ClipboardItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	        this.createdAt = source["createdAt"];
	    }
	}

}

