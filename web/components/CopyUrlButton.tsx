"use client";

import React, { useState } from "react";
import { Button, Tooltip } from "@mui/material";
import { Check, ContentCopy } from "@mui/icons-material";

export const CopyUrlButton = () => {
    const [hasCopied, setHasCopied] = useState(false);

    const copy = async () => {
        await navigator.clipboard.writeText(window.location.href);
        setHasCopied(true);

        setTimeout(() => setHasCopied(false), 2000);
        // TODO: Snackbar message
    };

    return (
        <Tooltip
            placement="top"
            title="Copy link to this page with all set URL parameters. Others opening the link should see the same view."
        >
            <Button
                onClick={copy}
                color="primary"
                variant="outlined"
                startIcon={hasCopied ? <Check /> : <ContentCopy />}
            >
                Copy link
            </Button>
        </Tooltip>
    );
};

export default CopyUrlButton;
