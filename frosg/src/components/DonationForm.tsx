import React from 'react';
import { Formik, Form, Field } from 'formik';
import * as yup from 'yup';

const validationSchema = yup.object({
    amount: yup.number().required('Amount is required').positive('Amount must be positive'),
});

const DonationForm: React.FC = () => {
    const handleSubmit = async (values: { amount: number }) => {
        // Implement payment process here with Clover API
    };

    return (
        <Formik
            initialValues={{ amount: 0 }}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
            >
                {({ errors, touched }) => (
                    <Form>
                        <Field type="number" name="amount" placeholder="Amount" />
                        {errors.amount && touched.amount ? <div>{errors.amount}</div> : null}
                        <button type='submit'>Donate</button>
                    </Form>
                )}
            </Formik>
    );
};

export default DonationForm