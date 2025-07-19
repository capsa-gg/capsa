import type * as React from "react";
import { useState } from "react";
import CircularProgress from "@mui/material/CircularProgress";
import SaveAltIcon from "@mui/icons-material/SaveAlt";
import { Button } from "@mui/material";
import { useSingleLogContext } from "@/context/SingleLogContext/SingleLogContext";
import { getJwtFromLocalStorage } from "@/data/jwt/localStorage";
import { BaseUrl } from "@/api/apibase";

const DownloadFullLogButton: React.FC = () => {
    const {
        metadata: { data: metadataData, isLoading },
    } = useSingleLogContext();

    const [isDownloading, setIsDownloading] = useState(false);

    // TODO: Should this be moved to a separate location?
    const handleDownloadLog = async () => {
        const id = metadataData?.logData.id;
        if (!id) return;

        try {
            setIsDownloading(true);
            const jwtData = getJwtFromLocalStorage();
            const jwt = jwtData?.token;
            if (!jwt) return;

            const baseUrl = await BaseUrl.GetBaseUrl();
            const unfilteredLogUrl = `${baseUrl}/v1/user/logs/${id}/log`;
            const response = await fetch(unfilteredLogUrl, {
                method: "GET",
                headers: {
                    accept: "application/json",
                    authorization: `Bearer ${jwt}`,
                },
            });

            // TODO: When Content-Length is supported, we can use response.body.getReader() to read the body per chunk and update a progress percentage
            const body = await response.text();

            // Create downloadable file
            const blob = new Blob([body], { type: "text/plain" });
            const url = URL.createObjectURL(blob);
            const a = document.createElement("a");
            a.href = url;
            a.download = `capsa-${id}.log`;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            URL.revokeObjectURL(url);
        } catch (e) {
            console.error("Error downloading full log", e);
        } finally {
            setIsDownloading(false);
        }
    };

    if (isLoading) return;
    return (
        <Button
            variant="outlined"
            startIcon={isDownloading ? <CircularProgress size={20} color="inherit" /> : <SaveAltIcon />}
            onClick={handleDownloadLog}
            disabled={isDownloading}
        >
            Full Log
        </Button>
    );
};

export default DownloadFullLogButton;
