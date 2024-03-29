# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GcsSourceResponse',
    'GlossaryInputConfigResponse',
    'GlossaryTermResponse',
    'GlossaryTermsPairResponse',
    'GlossaryTermsSetResponse',
    'LanguageCodePairResponse',
    'LanguageCodesSetResponse',
]

@pulumi.output_type
class GcsSourceResponse(dict):
    """
    The Google Cloud Storage location for the input content.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "inputUri":
            suggest = "input_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GcsSourceResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GcsSourceResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GcsSourceResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 input_uri: str):
        """
        The Google Cloud Storage location for the input content.
        :param str input_uri: Source data URI. For example, `gs://my_bucket/my_object`.
        """
        pulumi.set(__self__, "input_uri", input_uri)

    @property
    @pulumi.getter(name="inputUri")
    def input_uri(self) -> str:
        """
        Source data URI. For example, `gs://my_bucket/my_object`.
        """
        return pulumi.get(self, "input_uri")


@pulumi.output_type
class GlossaryInputConfigResponse(dict):
    """
    Input configuration for glossaries.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "gcsSource":
            suggest = "gcs_source"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GlossaryInputConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GlossaryInputConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GlossaryInputConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 gcs_source: 'outputs.GcsSourceResponse'):
        """
        Input configuration for glossaries.
        :param 'GcsSourceResponse' gcs_source: Google Cloud Storage location of glossary data. File format is determined based on the filename extension. API returns [google.rpc.Code.INVALID_ARGUMENT] for unsupported URI-s and file formats. Wildcards are not allowed. This must be a single file in one of the following formats: For unidirectional glossaries: - TSV/CSV (`.tsv`/`.csv`): Two column file, tab- or comma-separated. The first column is source text. The second column is target text. No headers in this file. The first row contains data and not column names. - TMX (`.tmx`): TMX file with parallel data defining source/target term pairs. For equivalent term sets glossaries: - CSV (`.csv`): Multi-column CSV file defining equivalent glossary terms in multiple languages. See documentation for more information - [glossaries](https://cloud.google.com/translate/docs/advanced/glossary).
        """
        pulumi.set(__self__, "gcs_source", gcs_source)

    @property
    @pulumi.getter(name="gcsSource")
    def gcs_source(self) -> 'outputs.GcsSourceResponse':
        """
        Google Cloud Storage location of glossary data. File format is determined based on the filename extension. API returns [google.rpc.Code.INVALID_ARGUMENT] for unsupported URI-s and file formats. Wildcards are not allowed. This must be a single file in one of the following formats: For unidirectional glossaries: - TSV/CSV (`.tsv`/`.csv`): Two column file, tab- or comma-separated. The first column is source text. The second column is target text. No headers in this file. The first row contains data and not column names. - TMX (`.tmx`): TMX file with parallel data defining source/target term pairs. For equivalent term sets glossaries: - CSV (`.csv`): Multi-column CSV file defining equivalent glossary terms in multiple languages. See documentation for more information - [glossaries](https://cloud.google.com/translate/docs/advanced/glossary).
        """
        return pulumi.get(self, "gcs_source")


@pulumi.output_type
class GlossaryTermResponse(dict):
    """
    Represents a single glossary term
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "languageCode":
            suggest = "language_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GlossaryTermResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GlossaryTermResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GlossaryTermResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 language_code: str,
                 text: str):
        """
        Represents a single glossary term
        :param str language_code: The language for this glossary term.
        :param str text: The text for the glossary term.
        """
        pulumi.set(__self__, "language_code", language_code)
        pulumi.set(__self__, "text", text)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> str:
        """
        The language for this glossary term.
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def text(self) -> str:
        """
        The text for the glossary term.
        """
        return pulumi.get(self, "text")


@pulumi.output_type
class GlossaryTermsPairResponse(dict):
    """
    Represents a single entry for an unidirectional glossary.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceTerm":
            suggest = "source_term"
        elif key == "targetTerm":
            suggest = "target_term"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GlossaryTermsPairResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GlossaryTermsPairResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GlossaryTermsPairResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_term: 'outputs.GlossaryTermResponse',
                 target_term: 'outputs.GlossaryTermResponse'):
        """
        Represents a single entry for an unidirectional glossary.
        :param 'GlossaryTermResponse' source_term: The source term is the term that will get match in the text,
        :param 'GlossaryTermResponse' target_term: The term that will replace the match source term.
        """
        pulumi.set(__self__, "source_term", source_term)
        pulumi.set(__self__, "target_term", target_term)

    @property
    @pulumi.getter(name="sourceTerm")
    def source_term(self) -> 'outputs.GlossaryTermResponse':
        """
        The source term is the term that will get match in the text,
        """
        return pulumi.get(self, "source_term")

    @property
    @pulumi.getter(name="targetTerm")
    def target_term(self) -> 'outputs.GlossaryTermResponse':
        """
        The term that will replace the match source term.
        """
        return pulumi.get(self, "target_term")


@pulumi.output_type
class GlossaryTermsSetResponse(dict):
    """
    Represents a single entry for an equivalent term set glossary. This is used for equivalent term sets where each term can be replaced by the other terms in the set.
    """
    def __init__(__self__, *,
                 terms: Sequence['outputs.GlossaryTermResponse']):
        """
        Represents a single entry for an equivalent term set glossary. This is used for equivalent term sets where each term can be replaced by the other terms in the set.
        :param Sequence['GlossaryTermResponse'] terms: Each term in the set represents a term that can be replaced by the other terms.
        """
        pulumi.set(__self__, "terms", terms)

    @property
    @pulumi.getter
    def terms(self) -> Sequence['outputs.GlossaryTermResponse']:
        """
        Each term in the set represents a term that can be replaced by the other terms.
        """
        return pulumi.get(self, "terms")


@pulumi.output_type
class LanguageCodePairResponse(dict):
    """
    Used with unidirectional glossaries.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceLanguageCode":
            suggest = "source_language_code"
        elif key == "targetLanguageCode":
            suggest = "target_language_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LanguageCodePairResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LanguageCodePairResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LanguageCodePairResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_language_code: str,
                 target_language_code: str):
        """
        Used with unidirectional glossaries.
        :param str source_language_code: The ISO-639 language code of the input text, for example, "en-US". Expected to be an exact match for GlossaryTerm.language_code.
        :param str target_language_code: The ISO-639 language code for translation output, for example, "zh-CN". Expected to be an exact match for GlossaryTerm.language_code.
        """
        pulumi.set(__self__, "source_language_code", source_language_code)
        pulumi.set(__self__, "target_language_code", target_language_code)

    @property
    @pulumi.getter(name="sourceLanguageCode")
    def source_language_code(self) -> str:
        """
        The ISO-639 language code of the input text, for example, "en-US". Expected to be an exact match for GlossaryTerm.language_code.
        """
        return pulumi.get(self, "source_language_code")

    @property
    @pulumi.getter(name="targetLanguageCode")
    def target_language_code(self) -> str:
        """
        The ISO-639 language code for translation output, for example, "zh-CN". Expected to be an exact match for GlossaryTerm.language_code.
        """
        return pulumi.get(self, "target_language_code")


@pulumi.output_type
class LanguageCodesSetResponse(dict):
    """
    Used with equivalent term set glossaries.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "languageCodes":
            suggest = "language_codes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LanguageCodesSetResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LanguageCodesSetResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LanguageCodesSetResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 language_codes: Sequence[str]):
        """
        Used with equivalent term set glossaries.
        :param Sequence[str] language_codes: The ISO-639 language code(s) for terms defined in the glossary. All entries are unique. The list contains at least two entries. Expected to be an exact match for GlossaryTerm.language_code.
        """
        pulumi.set(__self__, "language_codes", language_codes)

    @property
    @pulumi.getter(name="languageCodes")
    def language_codes(self) -> Sequence[str]:
        """
        The ISO-639 language code(s) for terms defined in the glossary. All entries are unique. The list contains at least two entries. Expected to be an exact match for GlossaryTerm.language_code.
        """
        return pulumi.get(self, "language_codes")


