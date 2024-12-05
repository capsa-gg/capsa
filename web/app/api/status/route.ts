export async function GET(): Promise<Response> {
    const healthCheck = {
        status: "ok",
        timestamp: new Date().toISOString(),
    };

    return Response.json(healthCheck, {
        status: 200,
        headers: {
            "Content-Type": "application/json",
            "Cache-Control": "no-store, no-cache, must-revalidate, proxy-revalidate",
        },
    });
}
