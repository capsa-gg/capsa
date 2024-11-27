import { formatDate } from "@/util/formatDate";
import { describe, it } from "@jest/globals";

describe("formatDateString", () => {
    it("formats strings correctly", () => {
        const input = "2024-11-14T17:08:23.317Z";
        const result = formatDate(input);
        const expected = "2024-11-14, 5:08:23 pm";

        expect(result).toEqual(expected);
    });

    it("does not crash on incorrect string input", () => {
        const input = "i am not correct";
        const result = formatDate(input);
        const expected = "Invalid date";

        expect(result).toEqual(expected);
    });

    it("does not crash on null input", () => {
        const input = null;
        const result = formatDate(input);
        const expected = "Invalid date";

        expect(result).toEqual(expected);
    });
});
