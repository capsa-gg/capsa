import moment from "moment/moment";

// biome-ignore lint/suspicious/noExplicitAny: this is desired here
export const formatDate = (str: any) => moment(str).utc().format("YYYY-MM-DD, h:mm:ss a");
