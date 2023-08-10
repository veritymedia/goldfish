export namespace main {
	
	export class ClipboardMessage {
	    status: string;
	    text: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ClipboardMessage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.text = source["text"];
	        this.message = source["message"];
	    }
	}

}

