"use client";

import React, { Suspense, useEffect } from "react";
import { TextField, Button, Link, Box } from "@mui/material";
import { useRouter, useSearchParams } from "next/navigation";
import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { useUserLogin } from "@/api/hooks";
import { removeJwtCookie, setJwtCookie } from "@/data/jwt/cookiesClient";
import { removeJwtFromLocalStorage, setJwtInLocalStorage } from "@/data/jwt/localStorage";
import { yupAuthValidation } from "@/types/api/validation";
import useUser from "@/context/UserContext";

interface FormInputLogin {
    email: string;
    password: string;
}

const validationSchemaLogin = yup.object({
    email: yupAuthValidation.email,
    password: yupAuthValidation.password,
});

const LoginPage: React.FC = () => {
    const { trigger, isMutating, error, data } = useUserLogin();
    const router = useRouter();
    const searchParams = useSearchParams();
    const user = useUser();
    const redirectUrl = searchParams.get("redirect");

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
    } = useForm<FormInputLogin>({
        resolver: yupResolver<FormInputLogin>(validationSchemaLogin),
        mode: "onChange",
    });

    /* eslint-disable react-hooks/exhaustive-deps */
    // TODO: callback functions

    // Remove jwt info if this page is reached
    useEffect(() => {
        if (redirectUrl) {
            removeJwtCookie();
            removeJwtFromLocalStorage();
        }
    }, []);

    // When we have the response, set values, redirect
    useEffect(() => {
        if (!error && data) {
            setJwtCookie(data);
            setJwtInLocalStorage(data);
            user.reloadInfoFromLocalStorage(); // This will update the signed in status in nav

            if (redirectUrl) {
                router.replace(redirectUrl);
            } else {
                router.refresh();
            }
        }
    }, [data]);

    const onSubmit: SubmitHandler<FormInputLogin> = formData => {
        trigger(formData);
    };

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
            <TextField
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email"
                autoComplete="email"
                autoFocus
                {...register("email")}
                error={Boolean(errors.email)}
                helperText={errors.email?.message}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                id="password"
                label="Password"
                type="password"
                autoComplete="current-password"
                {...register("password")}
                error={Boolean(errors.password)}
                helperText={errors.password?.message}
            />
            <Button disabled={!isValid || isMutating} type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
                Login
            </Button>
            <Link href="/auth/password-forgot" variant="body2" color="textSecondary" display="block" align="center">
                Forgot your password?
            </Link>
        </Box>
    );
};

const SuspenseLoginPage = () => (
    <Suspense>
        <LoginPage />
    </Suspense>
);

export default SuspenseLoginPage;
