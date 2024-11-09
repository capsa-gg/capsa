import * as jose from "jose";
import { type JwkJwt, type JwkResponse, JwkResponseSchema } from "@/types/api/jwk-jwt";

const tokenId = "capsa-server-jwk";
const algorithm = "RS256";
const jwkUri = ".well-known/jwks.json";
const expectedAudience = "capsa-user";
const expectedIssuer = "capsa-server";

// Class used to cache values
// Webapp needs to be restarted if the key changes
export default class JwtValidator {
    private static _jwk: JwkJwt | null = null;
    private static _pubKey: string | null = null;

    // Can throw errors
    private static async getJwk(): Promise<JwkJwt> {
        if (this._jwk) return this._jwk;

        const jwkEndpoint = `${process.env.NEXT_PUBLIC_SERVER_URL}/${jwkUri}`;
        console.log("jwkEndpoint", jwkEndpoint);

        // Load and process jwk, find the key we need
        const jwkRes = await fetch(jwkEndpoint);
        const jwkResJson = await jwkRes.json();
        const jwkResJsonParsed = JwkResponseSchema.parse(jwkResJson);
        const jwks: JwkResponse["keys"] = jwkResJsonParsed.keys;
        const jwk = jwks.find(j => j.kid === tokenId);
        if (!jwk) {
            console.error("jwk not found with key", tokenId);
            throw new Error("jwk not found");
        }

        this._jwk = jwk;
        return jwk;
    }

    public static async validateJwt(jwt: string): Promise<boolean> {
        const jwk = await this.getJwk();

        try {
            // Throws if validation fails
            await jose.importJWK(jwk, algorithm);
            const validated = await jose.jwtVerify(jwt, jwk, {
                algorithms: [algorithm],
                audience: expectedAudience,
                issuer: expectedIssuer,
            });

            // Sanity check
            return validated.payload.iss === expectedIssuer && validated.payload.aud === expectedAudience;
        } catch (e) {
            console.warn("decoding jwt failed", e);
            return false;
        }
    }
}
