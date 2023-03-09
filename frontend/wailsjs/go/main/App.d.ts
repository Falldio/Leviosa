// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {db} from '../models';

export function AddRSSFeed(arg1:string):Promise<db.Feed>;

export function ExportFeeds():Promise<void>;

export function FetchUpdates(arg1:number):Promise<void>;

export function FetchUpdatesForAllFeeds():Promise<void>;

export function GetFeeds():Promise<Array<db.Feed>>;

export function GetPost(arg1:number):Promise<db.Post>;

export function GetPosts(arg1:number):Promise<Array<db.Post>>;

export function ImportFeeds():Promise<Array<db.Feed>>;
