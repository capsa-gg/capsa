"use client";

import { ErrorSnackbarItem } from "@/containers/ErrorSnackbar/ErrorSnackbar.components";
import { useErrors } from "@/context/ErrorContext";
import type React from "react";

const ErrorSnackbar: React.FC = () => {
    const { errors } = useErrors();

    return errors.map(e => <ErrorSnackbarItem key={e.title} err={e} />);
};

export default ErrorSnackbar;
