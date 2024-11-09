"use client";

import React, { useEffect, useState } from "react";
import { useErrors } from "@/context/ErrorContext";
import { ErrorSnackbarItem } from "@/containers/ErrorSnackbar/ErrorSnackbar.components";

const ErrorSnackbar: React.FC = () => {
    const { errors } = useErrors();

    return errors.map(e => <ErrorSnackbarItem key={e.title} err={e} />);
};

export default ErrorSnackbar;
