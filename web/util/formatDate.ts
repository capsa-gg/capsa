// TODO: write unit tests

import moment from "moment/moment";

export const formatDate = (str: any) => moment(str).utc().format("YYYY-MM-DD, h:mm:ss a");
