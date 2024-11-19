import { describe, it } from "@jest/globals";
import { formatDateString } from "@/util/formatDateString";

describe("formatDateString", () => {
    it("formats strings correctly", () => {
        const input = "2024-11-14T17:08:23.317Z";
        const result = formatDateString(input);
        const expected = "2024-11-14, 6:08:23 pm";

        expect(result).toEqual(expected);
    });

    it("does not crash on incorrect input", () => {
        const input = "";
        const result = formatDateString(input);
        const expected = "Invalid date";

        expect(result).toEqual(expected);
    });
});
