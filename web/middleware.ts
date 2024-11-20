import { NextRequest, NextResponse } from "next/server";
import JwtValidator from "@/server/jwt";
import { deleteJwtCookie, getJwtCookieFromRequest } from "@/data/jwt/cookiesServer";

// This middleware checks if users are logged in.
// This is not implemented for security reasons, as that is done on the server.
// It simply serves as a way to make sure users don't get 4xx errors.
export async function middleware(req: NextRequest) {
    const token = await getJwtCookieFromRequest(req);
    const isAuthenticated = token ? await JwtValidator.validateJwt(token) : false;
    const isAuthRoute = req.nextUrl.pathname.startsWith("/auth");

    // Logged-in users accessing login routes should be redirected to the homepage
    if (isAuthenticated && isAuthRoute) {
        return NextResponse.redirect(new URL("/", req.url));
    }

    // Sent users that are not logged requesting non-auth pages in to the login page
    if (!isAuthenticated && !isAuthRoute) {
        await deleteJwtCookie(req);

        const url = new URL("/auth/login", req.url);
        url.searchParams.set("redirect", req.url);
        return NextResponse.redirect(url);
    }

    return NextResponse.next();
}

export const config = {
    matcher: [
        /*
         * Match all request paths except for the ones starting with:
         * - api (API routes)
         * - _next/static (static files)
         * - _next/image (image optimization files)
         * - favicon.ico, sitemap.xml, robots.txt (metadata files)
         */
        "/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)",
    ],
};
