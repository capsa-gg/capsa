import { deleteJwtCookie, getJwtCookieFromRequest } from "@/data/jwt/cookiesServer";
import JwtValidator from "@/server/jwt";
import { type NextRequest, NextResponse } from "next/server";

// This middleware checks if users are logged in.
// This is not implemented for security reasons, as that is done on the server.
// It simply serves as a way to make sure users don't get 4xx errors.
export async function middleware(req: NextRequest) {
    const { pathname, search, origin, basePath } = req.nextUrl;
    const path = `${basePath}${pathname === "/" ? "" : pathname}${search}`;
    const logBase = `[middleware][${req.method} ${path}]:`;

    const token = await getJwtCookieFromRequest(req);
    const isAuthenticated = token ? await JwtValidator.ValidateJwt(token) : false;
    const isAuthRoute = req.nextUrl.pathname.startsWith("/auth");

    console.log(logBase, "received request, isAuthenticated", isAuthenticated, "| isAuthRoute", isAuthRoute);

    // Logged-in users accessing login routes should be redirected to the homepage
    if (isAuthenticated && isAuthRoute) {
        console.log(logBase, "redirecting logged in user from auth route to", req.url);

        return NextResponse.redirect(new URL("/", req.url));
    }

    // Sent users that are not logged requesting non-auth pages in to the login page
    if (!isAuthenticated && !isAuthRoute) {
        await deleteJwtCookie(req);

        const signInUrl = new URL(`${basePath}/auth/login`, origin);

        if (path !== "") {
            signInUrl.searchParams.set("redirect", path);
        }

        console.log(logBase, "redirecting non logged in user to", signInUrl.toString());

        return NextResponse.redirect(signInUrl);
    }

    console.log(logBase, "continue executing request");

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
