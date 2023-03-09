export namespace db {
	
	export class Feed {
	    id: number;
	    url: string;
	    title: string;
	    description: string;
	    image: string;
	    updated: number;
	    created: number;
	    pinned: number;
	
	    static createFrom(source: any = {}) {
	        return new Feed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.url = source["url"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.image = source["image"];
	        this.updated = source["updated"];
	        this.created = source["created"];
	        this.pinned = source["pinned"];
	    }
	}
	export class Post {
	    id: number;
	    feed_id: number;
	    title: string;
	    description: string;
	    url: string;
	    content: string;
	    updated: number;
	    author: string;
	    read: boolean;
	    starred: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Post(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.feed_id = source["feed_id"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.url = source["url"];
	        this.content = source["content"];
	        this.updated = source["updated"];
	        this.author = source["author"];
	        this.read = source["read"];
	        this.starred = source["starred"];
	    }
	}

}

