"use client";

import { useUserPasswordResetComplete } from "@/api/hooks";
import { yupAuthValidation } from "@/types/api/validation";
import { yupResolver } from "@hookform/resolvers/yup";
import { Box, Button, Link, TextField, Typography } from "@mui/material";
import { useRouter, useSearchParams } from "next/navigation";
import type React from "react";
import { Suspense, useEffect, useState } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";

interface FormInputPasswordReset {
    resetToken: string;
    password: string;
}

const validationSchemaPasswordReset = yup.object({
    resetToken: yupAuthValidation.resetToken,
    password: yupAuthValidation.password,
});

const PasswordReset: React.FC = () => {
    const { trigger, isMutating, error } = useUserPasswordResetComplete();
    const [isSent, setIsSent] = useState(false);
    const router = useRouter();
    const searchParams = useSearchParams();

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
        setValue,
    } = useForm<FormInputPasswordReset>({
        resolver: yupResolver(validationSchemaPasswordReset),
        mode: "onChange",
    });

    // Remove jwt info if this page is reached
    useEffect(() => {
        const resetTokenParam = searchParams.get("resetToken");
        if (resetTokenParam) {
            setValue("resetToken", resetTokenParam);
        }
    }, [searchParams, setValue]);

    const onSubmit: SubmitHandler<FormInputPasswordReset> = async data => {
        await trigger(data);
        setIsSent(true);
    };

    if (isSent && !error && !isMutating) {
        return (
            <Box sx={{ display: "flex", alignItems: "center", flexDirection: "column" }}>
                <Typography align="center" mb={4}>
                    Password is set, please log in to continue
                </Typography>
                <Button variant="contained" onClick={() => router.push("/auth/login")}>
                    Login
                </Button>
            </Box>
        );
    }

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
            <TextField
                margin="normal"
                required
                fullWidth
                id="resetToken"
                label="Reset token"
                type="text"
                {...register("resetToken")}
                error={Boolean(errors.resetToken)}
                helperText={errors.resetToken?.message}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                id="password"
                label="New password"
                type="password"
                autoComplete="new-password"
                {...register("password")}
                error={Boolean(errors.password)}
                helperText={errors.password?.message}
            />
            <Button disabled={!isValid || isMutating} type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
                Set new password
            </Button>
            <Link href="/auth/login" variant="body2" color="textSecondary" display="block" align="center">
                Log in instead?
            </Link>
        </Box>
    );
};

const SuspensePasswordReset = () => (
    <Suspense>
        <PasswordReset />
    </Suspense>
);

export default SuspensePasswordReset;
