// TODO: write unit tests

import moment from "moment/moment";

export const formatDateString = (str: string) => moment(str).utc().format("YYYY-MM-DD, h:mm:ss a");
