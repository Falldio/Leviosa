import { Store } from "vuex";
import { main } from "../wailsjs/go/models";

declare module "@vue/runtime-core" {
    interface State {
        errorAlertInfo: {
            alert: boolean,
            title: string,
            message: string
        },
        feeds: main.Feed[]
    }
    interface ComponentCustomProperties {
        $store: Store<State>;
    }
}